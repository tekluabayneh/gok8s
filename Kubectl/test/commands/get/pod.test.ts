import { expect } from 'chai'
import { runCommand } from '@oclif/test'
import nock from 'nock'
import { fakePodsForTest } from '../../fixtures/pod.js'
import creaetApi from '../../../src/client/create_api.js'

describe('get pods', () => {

  afterEach(() => {
    nock.cleanAll()
  })

  it('gets and renders pods', async () => {
    // HOT: 
    // since create api is using url from kubectl this test fail in github action which need much more isolation than this 

    const api = await creaetApi()
    const scope = nock(api?.getUri()!)
      .get('/api/v1/namespaces/default/pods')
      .reply(200, fakePodsForTest)

    const { stdout } = await runCommand('get pods')
    expect(scope.isDone()).to.equal(true)
    expect(stdout).to.include('nginx')
    expect(stdout).to.include('redis')
  })
})


