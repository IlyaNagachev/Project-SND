import axios, {AxiosInstance, AxiosRequestConfig, AxiosResponse} from "axios";
import {config} from "./config";

class Network {
	private userApi : AxiosInstance = axios.create({
		baseURL: config.getProdUrl(),
		timeout: 1000,
		headers: {"Content-Type" : "application/json"}
	});

	UserPost<T>(url : string, payload : Record<string,any>){
		return this.sendPost<T>(url, payload)
	}

	private sendPost<T>(url : string, payload : Record<string,any>):Promise<T>{
		const r_cfg : AxiosRequestConfig = {
			timeout : 30000,
			withCredentials : true
		}
		return new Promise<T>((resolve, reject) => {
			this.userApi.post(url, payload, r_cfg)
				.then((response : AxiosResponse) => {
					resolve(response.data)
				})
				.catch((err) => {
					reject(err)
				})
		})
	}

}
const network : Network = new Network()
export default network
