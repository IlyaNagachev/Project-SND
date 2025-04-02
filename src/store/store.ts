import {SucceedResponse, User} from "../models/user";
import {defineStore} from "pinia";
import network from "../network/network";
import {Product, Requests} from "../models/requests";

interface state {
	Products : Product[],
	SelectedProducts : Product
}

export const ProductStore = defineStore("users", {
	state : () : state => ({
		Products : <Product[]>{},
		SelectedProducts: <Product>{}
	}),
	actions : {
		Registration(login : string, password : string) : Promise<SucceedResponse>{
			return new Promise<SucceedResponse>((resolve, reject) => {
				network.UserPost<SucceedResponse>(Requests.Registration, {
					login : login,
					password : password
				})
					.then((r) => {
						resolve(r)
						return(r)
					})
					.catch((err) => {
						reject(err)
					})
			})
		},
		GetProducts() : Promise<Product[]>{
			return new Promise<Product[]>((resolve, reject) => {
				network.UserPost<Product[]>(Requests.GetProduct, {})
					.then((r) => {
						resolve(r)
						this.Products = r
						return(r)
					})
					.catch((err) => {
						reject(err)
					})
			})
		},

		GetProduct(uuid : number) : Promise<Product>{
			return new Promise<Product>((resolve, reject) => {
				network.UserPost<Product>(Requests.GetProductById, {uuid : uuid})
					.then((r) => {
						resolve(r)
						return(r)
					})
					.catch((err) => {
						reject(err)
					})
			})
		}
	}
})

