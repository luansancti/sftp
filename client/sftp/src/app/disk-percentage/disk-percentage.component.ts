import { Component, OnInit } from '@angular/core';
import {ThemePalette} from '@angular/material/core';
import {ProgressSpinnerMode} from '@angular/material/progress-spinner';
import { DiskPercentageService } from './disk-percentage.service';
import {DiskUsage} from '../models/user'

@Component({
  selector: 'app-disk-percentage',
  templateUrl: './disk-percentage.component.html',
  styleUrls: ['./disk-percentage.component.scss']
})
export class DiskPercentageComponent implements OnInit {

  panelOpenState = false;
  data: DiskUsage[]
  constructor(private _diskService: DiskPercentageService) { }

  ngOnInit(): void {
    this._diskService.GetUsersConnected().subscribe(x => {
      if(x.Success) {
        this.data = x.Data
        console.log(this.data)
      }
    })
  }

  color: ThemePalette = 'primary';
  mode: ProgressSpinnerMode = 'determinate';
  value = 50;

}
