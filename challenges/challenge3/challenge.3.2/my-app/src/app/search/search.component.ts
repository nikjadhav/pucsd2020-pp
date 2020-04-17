import { Component, OnInit } from '@angular/core';
import { HttpErrorResponse } from '@angular/common/http';
import { UserService} from '../user.service';
import { Routes, RouterModule, Router } from '@angular/router';
import { FormBuilder, FormGroup, FormControl, Validators } from '@angular/forms';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})  
export class SearchComponent implements OnInit {
  data:any;
  Id: string="";
  users=[];
  fname="";
  lname="";
  email="";
  contact="";
  pass="";
  
  
  constructor(
   private formBuilder:FormBuilder,
   private userservice:UserService,
   private router:Router,
   private toastr:ToastrService,
  ) { }
  ngOnInit():void{  
  }
Reset()
{
  this.Id="";
}
Search(){

  this.userservice.get_userbyid(this.Id).subscribe
     (
       (data:any[])=>
       {
         
        console.log("Data",data);
        var stringData = '[' + JSON.stringify(data) + ']'
        var parseData = JSON.parse(stringData);
        console.log("parse",parseData[0][0]);
        this.users = parseData;
        console.log("In",this.users);     
       },
       (err:any[]) =>
       {

       }
       
      
     )
     console.log("usr",this.users);
     this.userservice.set_display_search(false);
    }
   
    Delete(id){
      this.toastr.success("DELETE user successfully","TITLE");
      this.userservice.del_user(id).subscribe
      (
        (response)=>
        {
          
            
            console.log(response); 
        },
        (error) =>
        {
            console.log(error);
        }
       
      )
      this.router.navigate(['/']); 
    }

     
  Update(){
    let update_data={
      "first_name":this.fname,"last_name":this.lname,"email":this.email,"contact_number":this.contact,"password":this.pass
    }
    this.userservice.update_user(this.Id,update_data).subscribe
    (
      (data:any[])=>
      {

      },
      (err:any[]) =>
      {
        
      }
    )
    this.router.navigate(['/']); //navigate to ho
  }



    is_display_search(){
      return this.userservice.get_display_search();
   }
   /*change display settings */
   update_display_setting(){  
     this.userservice.set_display_search(true);
     this.userservice.set_display_update(false);
     /*setting up already filled values*/
     this.fname=this.users[0].data.first_name;
     this.lname=this.users[0].data.last_name;
     this.email=this.users[0].data.email;
     this.pass=this.users[0].data.password;
     this.contact=this.users[0].data.contact_number;
  
   }
   is_display_update(){
     return this.userservice.get_display_update();
   }
  


}