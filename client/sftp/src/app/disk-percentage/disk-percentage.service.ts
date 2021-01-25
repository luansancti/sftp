import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import {GatewayService} from '../helper/gateway.service'
import {ResponseDiskPercentage} from '../models/user'

@Injectable({
  providedIn: 'root'
})
export class DiskPercentageService {

  constructor(private gateway: GatewayService) { }

  public GetUsersConnected(): Observable<ResponseDiskPercentage> {
    return this.gateway.GET("percentagedisk")
  }
}
