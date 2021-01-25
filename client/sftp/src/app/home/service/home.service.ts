import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import {GatewayService} from '../../helper/gateway.service'
import {ResponseUserList, ResponseGeneric, UserAdd, ResponseKey} from '../../models/user'



@Injectable({
  providedIn: 'root'
})
export class HomeService {

  constructor(private gateway: GatewayService) { }

  public GetListUser(): Observable<ResponseUserList> {
    return this.gateway.GET("listusers")
  }

  public Fix(content: any): Observable<ResponseGeneric> {
    return this.gateway.POST("fixpermissionuser",content)
  }

  public DeleteUser(content: UserAdd): Observable<ResponseGeneric> {
    return this.gateway.POST("deleteuser",content)
  }

  public DownlaodKey(content: UserAdd): Observable<ResponseKey> {
    return this.gateway.POST("downlaodkey",content)
  }

  public ChangePassword(content: UserAdd): Observable<ResponseKey> {
    return this.gateway.POST("changepassword",content)
  }

  public ChangeExpiration(content: UserAdd): Observable<ResponseKey> {
    return this.gateway.POST("changeexpiration",content)
  }


}
