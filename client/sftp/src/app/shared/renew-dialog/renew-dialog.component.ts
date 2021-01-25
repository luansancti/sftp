import { Component, Input, OnInit } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import * as moment from 'moment';

@Component({
  selector: 'app-renew-dialog',
  templateUrl: './renew-dialog.component.html',
  styleUrls: ['./renew-dialog.component.scss']
})
export class RenewDialogComponent implements OnInit {

  @Input() title: string;
  @Input() message: string;
  @Input() btnOkText: string;
  @Input() btnCancelText: string;
  value: number;
  expirationDate: string
  showP: boolean

  constructor(private activeModal: NgbActiveModal) { }

  ngOnInit() {
  }

  public decline() {
    this.activeModal.close(false);
  }

  public accept() {
    this.activeModal.close(this.value);
  }

  public dismiss() {
    this.activeModal.dismiss();
  }

  public counter(i: number) {
    return new Array(i);
  }

  public selectOption(id: number) {
    this.value = id;
  }

  getWhenExpired(days) {
    console.log(days)
    if (days == 1) {
      this.showP = false
      return this.expirationDate = ""
    }
    this.showP = true
    var today = new Date();
    var tomorrow = new Date();
    tomorrow.setDate(today.getDate()+Number(days));
    tomorrow.setHours(0,0,0)
    this.expirationDate = moment(tomorrow.toString()).format('YYYY/MM/DD HH:mm')
  }

}
