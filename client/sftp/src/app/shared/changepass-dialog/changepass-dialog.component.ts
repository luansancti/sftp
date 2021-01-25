import { Component, OnInit, Input } from '@angular/core';
import { NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import {Helper} from '../../helper/helper'
import { Clipboard } from '@angular/cdk/clipboard';
import {FormControl, FormGroupDirective, FormGroup, NgForm, Validators} from '@angular/forms';


@Component({
  selector: 'app-changepass-dialog',
  templateUrl: './changepass-dialog.component.html',
  styleUrls: ['./changepass-dialog.component.scss']
})
export class ChangepassDialogComponent implements OnInit {

  @Input() title: string;
  @Input() message: string;
  @Input() btnOkText: string;
  @Input() btnCancelText: string;
  strength: number;

  constructor(private activeModal: NgbActiveModal, 
    private helper: Helper,
    private clipboard: Clipboard,) { }

    myForm = new FormGroup({
      password: new FormControl('', [
        
        Validators.minLength(8)
      ]),  
    });

  ngOnInit() {
  }

  public decline() {
    this.activeModal.close(false);
  }

  public accept() {
    this.activeModal.close(this.myForm.get('password').value);
  }

  public dismiss() {
    this.activeModal.dismiss();
  }

  public counter(i: number) {
    return new Array(i);
  }

  onStrengthChanged(strength: number) {
    this.strength = strength
  }

  randomPassword() {
    let password = this.helper.generatePassword(12, true, true, true)
    this.myForm.controls['password'].setValue(password)
    this.clipboard.copy(password)
  }

}
