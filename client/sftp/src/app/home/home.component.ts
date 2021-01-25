import { AfterViewInit, Component, OnInit, ViewChild } from '@angular/core';
import {HomeService} from './service/home.service'
import {MatPaginator} from '@angular/material/paginator';
import {MatTableDataSource} from '@angular/material/table';
import {UserAdd, UserDetails} from '../models/user'
import { animate, state, style, transition, trigger } from '@angular/animations';
import {Helper} from '../helper/helper'
import * as moment from 'moment';
import {MatSort} from '@angular/material/sort';
import {MatSnackBar} from '@angular/material/snack-bar';
import { ConfirmationDialogService } from '../shared/confirmation-dialog/confirmation-dialog-service.service';
import { RenewDialogService } from '../shared/renew-dialog/renew-dialog.service';
import {ChangepassDialogService} from '../shared/changepass-dialog/changepass-dialog.service'



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


export class HomeComponent implements AfterViewInit {

  constructor(private homeService: HomeService, 
    public helper: Helper, 
    private _snackBar: MatSnackBar,
    private confirmationDialogService: ConfirmationDialogService,
    private renewDialog: RenewDialogService,
    private changeDialog: ChangepassDialogService
    ) { }
 
  displayedColumns: string[] = ['UserName', 'Expiration', 'Size'];
  dataSource: MatTableDataSource<UserDetails>
  panelOpenState = false;
  hide = true;

  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild(MatSort) sort: MatSort;

  isExpansionDetailRow = (i: number, row: Object) => row.hasOwnProperty('detailRow'); 
  
  writeContents(content, fileName, contentType) {
    var a = document.createElement('a');
    var file = new Blob([content], {type: contentType});
    a.href = URL.createObjectURL(file);
    a.download = fileName;
    a.click();
  }

  public formatHour(dateSent){
    if (dateSent == "Never") {
      return dateSent
    }
    var d1 = moment(new Date(Date.now()).toUTCString());
    var d2 = moment(new Date(dateSent * 1000).toUTCString());
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
    this.loadUsers()
  }

  loadUsers() {
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

  changePassword(element: UserDetails) {
    this.changeDialog.confirm('Change Password','')
    .then(x => {

      if(x) {
        let useradd = new UserAdd()
        useradd.User = element.UserName
        useradd.Password = x.toString()
        this.homeService.ChangePassword(useradd)
        .subscribe(content => {
          this._snackBar.open(content.Message, 'OK', {
            duration: 2000,
            horizontalPosition: "center",
            verticalPosition: "top",
          });
        })
      }
    })
  }

  renewUser(element: UserDetails) {
    this.renewDialog.confirm('Renew User', 'Select in days do you want renew')
    .then(x => {
      console.log(x)
      if(x) {
        console.log("cai aqui")
        let useradd = new UserAdd()
        useradd.User = element.UserName
        useradd.Expiration = Number(x.toString())
        this.homeService.ChangeExpiration(useradd)
        .subscribe(x => {
          if(x.Success) {
            this._snackBar.open(x.Message, 'OK', {
              duration: 2000,
              horizontalPosition: "center",
              verticalPosition: "top",
            });
            this.loadUsers()
          } else {
            this._snackBar.open(x.Message, 'OK', {
              duration: 2000,
              horizontalPosition: "center",
              verticalPosition: "top",
            });
          }
        })
      }
    })
  }

  downloadKey(element: UserDetails) {
    let userAdd = new UserAdd()
    userAdd.User = element.UserName
    this.homeService.DownlaodKey(userAdd).subscribe(content => {
      if(content.Success) {

        this.downloadFile(content.Data, `${element.UserName}.rsa`, 'text/plain')
        this._snackBar.open(content.Message, 'OK', {
          duration: 2000,
          horizontalPosition: "center",
          verticalPosition: "top",
        });
      }
      
    })
  }


  fixPermission(element: UserDetails) {
    let userAdd = new UserAdd()
    userAdd.User = element.UserName
    this.homeService.Fix(userAdd)
    .subscribe(x => {
      if(x.Success) {
        this._snackBar.open(x.Message, 'OK', {
          duration: 2000,
          horizontalPosition: "center",
          verticalPosition: "top",
        });
      } else {
        this._snackBar.open(x.Message, 'OK', {
          duration: 2000,
          horizontalPosition: "center",
          verticalPosition: "top",
        });
      }
    })

  }

  downloadFile(content: any, fileName: string, contentType) {
    var a = document.createElement('a');
    var file = new Blob([content], {type: contentType});
    a.href = URL.createObjectURL(file);
    a.download = fileName;
    a.click();
  }

  deleteUser(element: UserDetails) {
    this.confirmationDialogService.confirm('Please confirm..', `Do you really want delete user ${element.UserName}?`)
    .then((confirmed) => {
      if(confirmed) {
        let useradd = new UserAdd()
      useradd.User = element.UserName
      this.homeService
      .DeleteUser(useradd)
      .subscribe(x => {
        if(x.Success) {
          this._snackBar.open(x.Message, 'OK', {
            duration: 2000,
            horizontalPosition: "center",
            verticalPosition: "top",
          });
          this.getNotification()
        } else {
          this._snackBar.open(x.Message, 'OK', {
            duration: 2000,
            horizontalPosition: "center",
            verticalPosition: "top",
          });
        }
      })
      }
      
    })
    
  }


  

  getNotification() {
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
