import { runCommand } from '@oclif/test'
import { expect } from 'chai'

import nock from 'nock'
describe('pods test', () => {
  const fakePods = {
    items: [
      {
        metadata: {
          name: 'nginx',
        },
        spec: {
          containers: [{}],
        },
        status: {
          phase: 'Running',
        },
      },
      {
        metadata: {
          name: 'redis',
        },
        spec: {
          containers: [{}],
        },
        status: {
          phase: 'Pending',
        },
      },
    ],
  }

  it('should print pods', async () => {
    const baseUrl = 'http://kubernetes.test'
    const nockRes = nock(baseUrl)
      .get(`/api/v1/namespaces/default/pods`)
      .reply(200, fakePods)

    const { stdout } = await runCommand(['mykubectl', 'get', 'pods'])
    console.log("nock", nockRes.isDone())
    console.log("stdout", stdout)
    expect(stdout).to.contain('nginx')
  })


  // when test run this contains is looking for string "pods" in the result but mykubectl is pring list of pods if they are appear to be in  k8s 
  // if not they are printed as only string that has 


})

