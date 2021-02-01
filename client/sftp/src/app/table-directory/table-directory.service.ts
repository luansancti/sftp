import { Injectable } from '@angular/core';
import {GatewayService} from '../helper/gateway.service'
import { Observable } from 'rxjs';
import {ReponseListDirectory} from '../models/user'

@Injectable({
  providedIn: 'root'
})

export class TableDirectoryService {

  constructor(private gateway: GatewayService) { }

  public ListDirectory(content: string): Observable<ReponseListDirectory> {
    return this.gateway.POST("listdirectory",content)
  }

}
