import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-disk-percentage',
  templateUrl: './disk-percentage.component.html',
  styleUrls: ['./disk-percentage.component.scss']
})
export class DiskPercentageComponent implements OnInit {

  panelOpenState = false;
  constructor() { }

  ngOnInit(): void {
  }

}
