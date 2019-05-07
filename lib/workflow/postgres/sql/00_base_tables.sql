CREATE TABLE workflows (
    id BIGSERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,  
    is_singleton BOOLEAN,
    recurrence TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    next_due TIMESTAMP NOT NULL DEFAULT NOW(),
    default_parameters JSON
);

CREATE TABLE workflow_instances (
    id BIGSERIAL PRIMARY KEY,
    workflow_id BIGINT NOT NULL REFERENCES workflows(id),
    parameters JSON,
    payload JSON,

    enqueued_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE tasks (
    id BIGSERIAL PRIMARY KEY,
    workflow_instance_id BIGINT NOT NULL REFERENCES workflow_instances(id),
    try_remaining INT NOT NULL DEFAULT 1,
    enqueued_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP NOT NULL DEFAULT NOW(),
    start_after   TIMESTAMP NOT NULL DEFAULT NOW(),
    task_name     TEXT NOT NULL,
    parameters    JSON
);

/* CREATE TABLE workflow_events (
    -- TaskComplete()
    event_type
    workflow_instance_id 
    task_result_id
);
 */
 
CREATE TYPE task_status AS ENUM('success', 'failed', 'abandoned');

CREATE TABLE tasks_results (
    id BIGSERIAL PRIMARY KEY,
    workflow_instance_id BIGINT NOT NULL REFERENCES workflow_instances(id),
    parameters   JSON,
    task_name    TEXT NOT NULL,
    enqueued_at  TIMESTAMP NOT NULL,
    completed_at TIMESTAMP NOT NULL DEFAULT NOW(),
    status       task_status,
    error        TEXT,
    result       JSON
);


-- Notification channels
-- 
-- workflow_task_new
-- workflow_task_complete
-- 

-- https://www.postgresql.org/docs/current/xfunc-sql.html
CREATE OR REPLACE FUNCTION enqueue_task(
    workflow_instance_id BIGINT, 
    try_remaining INT,
    start_after TIMESTAMP,
    task_name TEXT,
    parameters JSON)
RETURNS VOID
AS $$
    INSERT INTO tasks(workflow_instance_id, try_remaining, start_after, task_name, parameters) 
        VALUES(workflow_instance_id, try_remaining, start_after, task_name, parameters);
    SELECT pg_notify('workflow_task_new', task_name);
$$ LANGUAGE SQL;

--https://stackoverflow.com/questions/16609724/using-current-time-in-utc-as-default-value-in-postgresql
CREATE OR REPLACE FUNCTION dequeue_task(task_name TEXT)
RETURNS RECORD
AS $$
    UPDATE tasks t1 SET try_remaining = try_remaining - 1, updated_at = NOW()
    WHERE t1.id = (
        SELECT t2.id FROM tasks t2
        WHERE t2.try_remaining > 0 AND start_after < NOW()
        ORDER BY t2.enqueued_at FOR UPDATE SKIP LOCKED LIMIT 1        
    ) RETURNING t1.id, t1.parameters
$$ LANGUAGE SQL;

-- https://www.postgresql.org/docs/9.1/queries-with.html
-- https://stackoverflow.com/questions/6560447/can-i-use-return-value-of-insert-returning-in-another-insert#6561437
CREATE OR REPLACE FUNCTION complete_task(tid BIGINT, status task_status, error text, result JSON)
RETURNS VOID
LANGUAGE SQL
AS $$
    WITH done_tasks AS (SELECT workflow_instance_id AS id FROM tasks where id = tid)
    SELECT pg_notify('workflow_task_complete', id::text) FROM done_tasks;
    WITH in_vals AS (SELECT
        tid as id, 
        status as status, 
        error as error,
        result as result)
    INSERT INTO tasks_results(workflow_instance_id, parameters, task_name, enqueued_at, status, error, result)
        (SELECT workflow_instance_id, parameters, task_name, enqueued_at, in_vals.status, in_vals.error, in_vals.result
         FROM tasks JOIN in_vals ON in_vals.id = tasks.id WHERE tasks.id = tid);
    DELETE FROM tasks WHERE id = tid 
$$;

INSERT INTO workflows(name) VALUES ('test_workflow');
INSERT INTO workflow_instances(workflow_id) VALUES (1);