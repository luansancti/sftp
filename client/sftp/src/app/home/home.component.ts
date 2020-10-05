import { AfterViewInit, Component, OnInit, ViewChild } from '@angular/core';
import {HomeService} from './service/home.service'
import {MatPaginator} from '@angular/material/paginator';
import {MatTableDataSource} from '@angular/material/table';
import {UserDetails} from '../models/user'
import { animate, state, style, transition, trigger } from '@angular/animations';
import {Helper} from '../helper/helper'
import * as moment from 'moment';
import {MatSort} from '@angular/material/sort';


@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({height: '0px', minHeight: '0', display: 'none'})),
      state('expanded', style({height: '*'})),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ],
})


export class HomeComponent implements OnInit, AfterViewInit {

  constructor(private homeService: HomeService, public helper: Helper) { }
 
  displayedColumns: string[] = ['UserName', 'Expiration', 'Owner', 'Size'];
  dataSource: MatTableDataSource<UserDetails>
  panelOpenState = false;
  hide = true;

  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild(MatSort) sort: MatSort;

  isExpansionDetailRow = (i: number, row: Object) => row.hasOwnProperty('detailRow'); 
  
  ngOnInit(): void {
    
  }


  public formatHour(dateSent){
    if (dateSent == "Never") {
      return dateSent
    }
    var d1 = moment(Date.now());
    var d2 = moment(dateSent);
    console.log(dateSent)
    var days = moment.duration(d2.diff(d1)).asDays();
    days = Math.floor(days)
    if(days < 0) {
      return "Expired"
    } else if (days == 0) {
      return "Today"
    } else {
      return days
    }
  }

  ngAfterViewInit() {
    this.homeService.GetListUser()
    .subscribe(x => {
      if(x.Success) {
        x.Data.map(x => {
          x.Expiration = this.formatHour(x.Expiration)
        })
        this.dataSource = new MatTableDataSource<UserDetails>(x.Data);
        
      } else {
        this.dataSource = new MatTableDataSource<UserDetails>([]);
      }
      this.dataSource.paginator = this.paginator;
      this.dataSource.sort = this.sort;
    })
  }

  getNotification(evt) {
    // Do something with the notification (evt) sent by the child!
    this.homeService.GetListUser()
    .subscribe(x => {
      if(x.Success) {
        x.Data.map(x => {
          x.Expiration = this.formatHour(x.Expiration)
        })
        this.dataSource = new MatTableDataSource<UserDetails>(x.Data);
        
      } else {
        this.dataSource = new MatTableDataSource<UserDetails>([]);
      }
      this.dataSource.paginator = this.paginator;
      this.dataSource.sort = this.sort;
    })
  }

  


}
