import { Component, OnInit, ViewChild } from '@angular/core';
import { animate, state, style, transition, trigger } from '@angular/animations';
import {DashboardService} from './dashboard.service'
import {MatPaginator} from '@angular/material/paginator';
import {MatTableDataSource} from '@angular/material/table';
import {MatSort} from '@angular/material/sort';
import { UserAdd } from '../models/user';
import {MatSnackBar} from '@angular/material/snack-bar';



@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({height: '0px', minHeight: '0', display: 'none'})),
      state('expanded', style({height: '*'})),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ],
})
export class DashboardComponent implements OnInit {

  panelOpenState = false;
  constructor(private _dashService: DashboardService, private _snackBar: MatSnackBar) { }

  dataSource: MatTableDataSource<string>
  displayedColumns: string[] = ['UserName', 'Action'];

  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild(MatSort) sort: MatSort;

  isExpansionDetailRow = (i: number, row: Object) => row.hasOwnProperty('detailRow'); 

  ngOnInit(): void {
    this.get_users()
  }

  get_users() {
    this._dashService.GetUsersConnected()
    .subscribe(content => {
      this.dataSource = new MatTableDataSource<string>(content.Data);
    })
  }

  link_off(element: string) {
    let add_user = new UserAdd()
    add_user.User = element
    this._dashService.Unlink_User(add_user)
    .subscribe(x => {
      if(x.Success) {
        this._snackBar.open(x.Message, 'OK', {
          duration: 2000,
          horizontalPosition: "center",
          verticalPosition: "top",
        });
        this.get_users()
      } else {
        this._snackBar.open(x.Message, 'OK', {
          duration: 2000,
          horizontalPosition: "center",
          verticalPosition: "top",
        });
      }
      
    })
  }

}
