export enum Requests {
	Registration = "user/reg",
	GetProduct = "user/get",
	GetProductById = "product/id"

}
export interface Product{
	uuid : string,
	class : string,
	description : string,
	img : string,
	name : string,
	price : number,
	specs : Record<string, string>
}
