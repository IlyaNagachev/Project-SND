class Config {
	private readonly prodHost : string
	private readonly devHost : string

	constructor() {
		this.prodHost = "http://localhost:8001"
		this.devHost = "http://localhost:8001"
	}

	getProdUrl() : string {return this.prodHost}
	getDevUrl() : string {return this.devHost}
}

const config : Config = new Config()

export {config}


