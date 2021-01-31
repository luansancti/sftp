import { Injectable } from '@angular/core';
import {GatewayService} from '../helper/gateway.service'
import { Observable } from 'rxjs';
import {ResponseGeneric} from '../models/user'

@Injectable({
  providedIn: 'root'
})

export class TableDirectoryService {

  constructor(private gateway: GatewayService) { }

  public DeleteUser(content: string): Observable<ResponseGeneric> {
    return this.gateway.POST("deleteuser",content)
  }

}
