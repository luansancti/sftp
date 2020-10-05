import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import {GatewayService} from '../../helper/gateway.service'
import {ResponseUserList} from '../../models/user'



@Injectable({
  providedIn: 'root'
})
export class HomeService {

  constructor(private gateway: GatewayService) { }

  public GetListUser(): Observable<ResponseUserList> {
    return this.gateway.GET("listusers")
  }
}
