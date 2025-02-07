class Config {
	private readonly prodHost : string
	private readonly devHost : string

	constructor() {
		this.prodHost = "http://localhost:4001"
		this.devHost = "http://localhost:4002"
	}

	getProdUrl() : string {return this.prodHost}
	getDevUrl() : string {return this.devHost}
}

const config : Config = new Config()

export {config}


