import { runCommand } from '@oclif/test'
import { expect } from 'chai'

describe('hello', () => {
  it('runs hello', async () => {
    const { stdout } = await runCommand('pods')

    expect(stdout).to.contain('hello')
  })
})
