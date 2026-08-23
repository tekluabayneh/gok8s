
import { Command, Flags } from '@oclif/core'

export default class nodes extends Command {
  static flags = {
    from: Flags.string({ char: 'f', description: 'Who is saying hello', required: true }),
  }

  async run(): Promise<void> {
    this.log(`well this work i get pood is been fetched...`)
  }
}
