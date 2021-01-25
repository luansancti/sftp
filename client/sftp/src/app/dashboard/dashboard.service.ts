import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import {GatewayService} from '../helper/gateway.service'
import {ResponseData, ResponseGeneric, UserAdd} from '../models/user'

@Injectable({
  providedIn: 'root'
})
export class DashboardService {

  constructor(private gateway: GatewayService) { }

  public GetUsersConnected(): Observable<ResponseData> {
    return this.gateway.GET("userslogged")
  }

  public Unlink_User(content: UserAdd): Observable<ResponseGeneric> {
    return this.gateway.POST("unlink_user", content)
  }

}
