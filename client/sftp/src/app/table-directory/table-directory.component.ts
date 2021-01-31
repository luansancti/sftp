import { Component, OnInit } from '@angular/core';
import { animate, state, style, transition, trigger } from '@angular/animations';
import {MatSnackBar} from '@angular/material/snack-bar';
import { MatTableDataSource } from '@angular/material/table';
import {ListDirectory} from '../models/user'

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

  constructor(private _snackBar: MatSnackBar,) { }

  displayedColumns: string[] = ['UserName', 'Expiration', 'Size'];
  dataSource: MatTableDataSource<ListDirectory>
  panelOpenState = false;
  hide = true;



  ngOnInit(): void {
  }

}
