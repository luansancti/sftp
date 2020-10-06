import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

import {environment} from '../../environments/environment'



@Injectable({
  providedIn: 'root'
})
export class GatewayService {

  private url = environment.apiEndpoint

  constructor(private http: HttpClient) { }

  public GET(path: string): Observable<any>{
      return this.http.get(`${this.url}/${path}`)
  }
  public POST(path: string, content: any): Observable<any> {
    let httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'application/json',
      })
    }
    return this.http.post(`${this.url}/${path}`, content, httpOptions)
  }
}
