import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators } from '@angular/forms';
import { Router } from "@angular/router";
import { UserService} from '../user.service';
import { MustMatch } from './must-match.validator'; 
import { ToastrService } from 'ngx-toastr';
@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  error_messages = {
    'fname': [
      { type: 'required', message: 'First Name is required.' },
    ],

    'lname': [
      { type: 'required', message: 'Last Name is required.' }
    ],

    'email': [
      { type: 'required', message: 'Email is required.' },
      { type: 'minlength', message: 'Email length.' },
      { type: 'maxlength', message: 'Email length.' },
      { type: 'pattern', message: 'please enter a valid email address.' }
    ],

    'pass': [
      { type: 'required', message: 'password is required.' },
      { type: 'minlength', message: 'password length.' },
      { type: 'maxlength', message: 'password length.' }
    ],
    'contact':[
      {type:'required',message:'contact is required'},
      {type:'pattern',message:'invalid number'},
    ],
  }

  registerForm: FormGroup;
  fname = new FormControl('', [Validators.required, Validators.maxLength(30)]);
  lname = new FormControl('', [Validators.required, Validators.maxLength(30)]);
  email = new FormControl('', [Validators.required, Validators.email, Validators.maxLength(40),Validators.pattern(new RegExp("^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$"))]);
  contact = new FormControl('', [Validators.required, Validators.pattern(new RegExp("[0-9 ]{10}"))])
  pass = new FormControl('', [Validators.required, Validators.minLength(3), Validators.maxLength(16)]);

  constructor(
    private formBuilder:FormBuilder,  
    private userservice:UserService,
    private toastr: ToastrService,

  ) {

   }
  
  ngOnInit(){
    this.createFormValidations();
    
  }
  onSubmit(){
    this.toastr.success("New User Added","TITLE");

    let userData = {
     
      "fname": this.registerForm.value.fname,
      "lname": this.registerForm.value.lname,
      "email": this.registerForm.value.email,
      "contact": this.registerForm.value.contact,
      "pass": this.registerForm.value.pass,
      
    };
    if(this.registerForm.valid){
      console.log("valid");
      this.userservice.post_data(userData).subscribe
      (
        (data:any[])=>
        {
          console.log(data);
          
         
        },
        (err:any[]) =>
        {
          console.log(err);
        }
      )     
    }
    else{
    console.log("invalid"); 
    }
     this.registerForm.reset();
  }

  createFormValidations() {
    this.registerForm = this.formBuilder.group({
      fname: this.fname,
      lname: this.lname,
      email: this.email,
      contact: this.contact,
      pass: this.pass,
      
    }, 
    );
    
  }

 }