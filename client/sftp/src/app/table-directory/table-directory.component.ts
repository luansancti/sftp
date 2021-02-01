import { Component, OnInit } from '@angular/core';
import { animate, state, style, transition, trigger } from '@angular/animations';
import {MatSnackBar} from '@angular/material/snack-bar';
import { MatTableDataSource } from '@angular/material/table';
import {ListDirectory} from '../models/user'
import { TableDirectoryService } from './table-directory.service'
import {Helper} from '../helper/helper'

@Component({
  selector: 'table-directory',
  templateUrl: './table-directory.component.html',
  styleUrls: ['./table-directory.component.scss'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({height: '0px', minHeight: '0', display: 'none'})),
      state('expanded', style({height: '*'})),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ],
})
export class TableDirectoryComponent implements OnInit {

  constructor(private _snackBar: MatSnackBar,private tableService: TableDirectoryService, public helper: Helper) { }

  displayedColumns: string[] = ['FileName', 'Size', 'LastModification'];
  dataSource: MatTableDataSource<ListDirectory>
  panelOpenState = false;
  hide = true;
  user = "igor"
  history = []
  show = false





  ngOnInit(): void {
    this.tableService.ListDirectory(this.user)
    .subscribe(x => {
      if (x.Success) {
        this.dataSource = new MatTableDataSource<ListDirectory>(x.Data);
        this.history.push(this.user)
      }else {
        this.dataSource = new MatTableDataSource<ListDirectory>([]);
      }
    })
    if (this.history.length > 1) {
      this.show = true
    } else {
      this.show = false
    }
    
  }

  returnFolder() {
    this.history.pop()
    this.tableService.ListDirectory(this.history.join("/"))
    .subscribe(x => {
      if (x.Success) {
        this.dataSource = new MatTableDataSource<ListDirectory>(x.Data); 
      }else {
        this.dataSource = new MatTableDataSource<ListDirectory>([]);
      }
    })
    if (this.history.length > 1) {
      this.show = true
    } else {
      this.show = false
    }
  }

  getNotification(locale) {
    this.history.push(locale)
    this.tableService.ListDirectory(this.history.join("/"))
    .subscribe(x => {
      if (x.Success) {
        this.dataSource = new MatTableDataSource<ListDirectory>(x.Data); 
      }else {
        this.dataSource = new MatTableDataSource<ListDirectory>([]);
      }
    })
    if (this.history.length > 1) {
      this.show = true
    } else {
      this.show = false
    }
  }




}
