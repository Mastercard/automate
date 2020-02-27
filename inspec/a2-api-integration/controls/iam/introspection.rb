require_relative '../../../constants'

# encoding: utf-8
# copyright: 2017, Chef Software, Inc.
# license: All rights reserved
title 'iam introspection'

control 'iam-introspection-1' do
  title 'iam introspection'
  desc 'these checks ensure that introspection works, enabling the UI to precalculate the permissions of the logged-in user'

  describe 'introspect all' do
    it 'returns what we expect' do
      resp = automate_api_request('/apis/iam/v2/introspect')
      expect(resp.http_status).to eq 200
    end
  end

  describe 'introspect some' do
    it 'returns what we expect' do
      resp = automate_api_request('/apis/iam/v2/introspect_some',
                                  http_method: 'POST',
                                  request_body:  {
                                    paths: [
                                      '/compliance/reporting/stats/summary',
                                      '/compliance/reporting/stats/failures',
                                      '/compliance/reporting/stats/trend',
                                    ]
                                  }.to_json)
      expect(resp.http_status).to eq 200
      expect(resp.parsed_response_body[:endpoints]).to eq(
        {
          '/compliance/reporting/stats/failures': {
            'get': false,
            'put': false,
            'post': true,
            'delete': false,
            'patch': false
          },
          '/compliance/reporting/stats/summary': {
            'get': false,
            'put': false,
            'post': true,
            'delete': false,
            'patch': false
          },
          '/compliance/reporting/stats/trend': {
            'get': false,
            'put': false,
            'post': true,
            'delete': false,
            'patch': false
          },
        }
      )
    end
  end

  describe 'introspect with parameterized endpoint' do
    it 'returns what we expect' do
      resp = automate_api_request('/apis/iam/v2/introspect',
                                  http_method: 'POST',
                                  request_body:  {
                                    path: '/iam/v2/policies/foo'
                                  }.to_json)
      # The parameter is "foo" because the matching path is '/iam/v2/policies/{id}'
      expect(resp.http_status).to eq 200
      expect(resp.parsed_response_body[:endpoints]).to eq(
        {
          '/iam/v2/policies/foo': {
            'get': true,
            'put': true,
            'post': false,
            'delete': true,
            'patch': false
          },
        }
      )
    end
  end
end