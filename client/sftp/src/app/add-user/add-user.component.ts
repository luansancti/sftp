import { Component, OnInit, Injectable, Output, EventEmitter } from '@angular/core';
import {FormControl, FormGroupDirective, FormGroup, NgForm, Validators} from '@angular/forms';
import { Clipboard } from '@angular/cdk/clipboard';
import {MatSnackBar} from '@angular/material/snack-bar';
import {AddUserService} from './add-user.service';
import { UserAdd } from '../models/user';


@Injectable({
  providedIn: 'root'
})
@Component({
  selector: 'app-add-user',
  templateUrl: './add-user.component.html',
  styleUrls: ['./add-user.component.scss']
})

export class AddUserComponent {
  @Output() notifyParent: EventEmitter<any> = new EventEmitter();

  
  constructor(
    private clipboard: Clipboard, 
    private _snackBar: MatSnackBar, 
    private _addUserService: AddUserService,
  ) { }

  panelOpenState: boolean = false;
  showDetails: boolean = false;
  hide: boolean = true;
  strength: number;
  checked: boolean = false;
  showP: boolean = false;
  expirationDate: string;


  myForm = new FormGroup({
    email: new FormControl('', [
      
    ]),
    password: new FormControl('', [
      
      Validators.minLength(8)
    ]),
    expiration: new FormControl('',[
      
    ]),
    key: new FormControl('',[
      
    ])
  });
  
  enablePass() {
    let key = this.myForm.get('key').value
    console.log(key)
    if( key == 'true') {
      return this.myForm.get('password').disable()
    } 
    
    this.myForm.get('password').enable()
    
    
  }

  onSubmit() {

    let username = this.myForm.get('email').value;
    let password = this.myForm.get('password').value;
    let key = this.myForm.get('key').value
    let expiration = this.myForm.get('expiration').value

    if(username == '' || password == '' || key == '' || expiration == '') {
      return 
    }
    let _addUser = new UserAdd()

    _addUser.User = username
    _addUser.Password = password
    _addUser.Expiration = Number(expiration)
    if (key == 'false') {
      this._addUserService
      .AddUser(_addUser)
      .subscribe(x => {
        if(x.Success) {
          this._snackBar.open(x.Message, 'End now', {
            duration: 2000,
            horizontalPosition: "center",
            verticalPosition: "top",
          });
          this.cleanForm()
          return this.notifyParent.emit(true)
        }
        this._snackBar.open(x.Message, 'End now', {
          duration: 2000,
          horizontalPosition: "center",
          verticalPosition: "top",
        });  
      })
    } else {
      this._addUserService
      .AddUserWithKey(_addUser)
      .subscribe(x => {
        if(x.Success) {
          this._snackBar.open(x.Message, 'OK', {
            duration: 2000,
            horizontalPosition: "center",
            verticalPosition: "top",
          });
          this.cleanForm()
          return this.notifyParent.emit(true)
        }
        this._snackBar.open(x.Message, 'OK', {
          duration: 2000,
          horizontalPosition: "center",
          verticalPosition: "top",
        });
      })
    }
  }

  cleanForm() {
    this.showP = false
    this.myForm.reset()
    Object.keys(this.myForm.controls).forEach(key =>{
      this.myForm.controls[key].setErrors(null)
   });
  }

  getWhenExpired(days) {
    if (days == -1) {
      return this.expirationDate = "Never"
    }
    this.showP = true
    var today = new Date();
    var tomorrow = new Date();
    tomorrow.setDate(today.getDate()+Number(days));
    tomorrow.setHours(0,0,0)
    this.expirationDate = tomorrow.toString()
  }

  doSomething() {
    let expirationValue = this.myForm.get('expiration').value
    if(expirationValue != '') {
      this.getWhenExpired(expirationValue)
    }
  }

  onStrengthChanged(strength: number) {
    this.strength = strength
  }


  randomPassword() {
    let password = this.generatePassword(12, true, true, true)
    this.myForm.controls['password'].setValue(password)
    this.clipboard.copy(password)
  }

  counter(i: number) {
    return new Array(i);
  }

  generatePassword(length, addUpper, addSymbols, addNums) {
    var lower = "abcdefghijklmnopqrstuvwxyz";
    var upper = addUpper ? lower.toUpperCase() : "";
    var nums = addNums ? "0123456789" : "";
    var symbols = addSymbols ? "!#$%&()*+,-.:;<=>?@_" : "";

    var all = lower + upper + nums + symbols;
    while (true) {
        var pass = "";
        for (var i=0; i<length; i++) {
            pass += all[Math.random() * all.length | 0];
        }

        // criteria:
        if (!/[a-z]/.test(pass)) continue; // lowercase is a must
        if (addUpper && !/[A-Z]/.test(pass)) continue; // check uppercase
        if (addSymbols && !/\W/.test(pass)) continue; // check symbols
        if (addNums && !/\d/.test(pass)) continue; // check nums

        return pass; // all good
    }
  }



}
