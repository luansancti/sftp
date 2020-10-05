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

export class AddUserComponent implements OnInit {
  @Output() notifyParent: EventEmitter<any> = new EventEmitter();

  panelOpenState = false
  showDetails: boolean = false
  hide = true
  strength: number
  checked = false
  password;


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

  constructor(
    private clipboard: Clipboard, 
    private _snackBar: MatSnackBar, 
    private _addUserService: AddUserService,
    ) { }

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
    if (!key) {
      this._addUserService
      .AddUser(_addUser)
      .subscribe(x => {
        if(x.Success) {
          this._snackBar.open(x.Message, 'End now', {
            duration: 2000,
            horizontalPosition: "right",
            verticalPosition: "top",
          });
          this.cleanForm()
          return this.notifyParent.emit(true)
        }
        this._snackBar.open(x.Message, 'End now', {
          duration: 2000,
          horizontalPosition: "right",
          verticalPosition: "top",
        });  
      })
    } else {
      this._addUserService
      .AddUserWithKey(_addUser)
      .subscribe(x => {
        if(x.Success) {
          this._snackBar.open(x.Message, 'End now', {
            duration: 2000,
            horizontalPosition: "right",
            verticalPosition: "top",
          });
          this.cleanForm()
          return this.notifyParent.emit(true)
        }
        this._snackBar.open(x.Message, 'End now', {
          duration: 2000,
          horizontalPosition: "right",
          verticalPosition: "top",
        });
      })
    }
  }

  ngOnInit(): void {
  }

  cleanForm() {
    
    this.myForm.reset()
    
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
    var symbols = addSymbols ? "!#$%&\'()*+,-.:;<=>?@[\\]^_`{|}~" : "";

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
