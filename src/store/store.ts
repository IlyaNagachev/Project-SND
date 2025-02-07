import {SucceedResponse, User} from "../models/user";
import {defineStore} from "pinia";
import network from "../network/network";
import {Requests} from "../models/requests";

interface state {
	User : User
}

export const UserStore = defineStore("users", {
	state : () : state => ({
		User : <User>{}
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
		}
	}
})

