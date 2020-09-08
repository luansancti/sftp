import { Injectable } from "@angular/core";
import { HttpClient } from '@angular/common/http';
import { User } from '../models/user'


@Injectable({
    providedIn: 'root'
})

export class Helper {
    constructor(private http: HttpClient) { }

    public GET(params: User) {
        this.http.get()
    }
}
