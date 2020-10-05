import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import {GatewayService} from '../helper/gateway.service'
import {UserAdd, ResponseGeneric} from '../models/user'



@Injectable({
  providedIn: 'root'
})
export class AddUserService {

  constructor(private gateway: GatewayService) { }

  public AddUser(user: UserAdd): Observable<ResponseGeneric> {
    return this.gateway.POST("createuser", JSON.stringify(user))
  }
  
  public AddUserWithKey(user: UserAdd): Observable<ResponseGeneric> {
    return this.gateway.POST("createuserwithkey", JSON.stringify(user))
  }
}
