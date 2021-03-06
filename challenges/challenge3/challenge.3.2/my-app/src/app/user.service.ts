import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpParams, HttpHeaders } from "@angular/common/http";
import { throwError } from 'rxjs';

@Injectable({
  providedIn: 'root'
  
})
export class UserService {
  is_display_search=true;/*if true then dont show search result table else show*/
  is_display_update=true;/*if true then dont show update form else show*/
  constructor(
    private httpClient: HttpClient
  ) { 
  }
  /*allow external api request and response*/
  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json','Access-Control-Allow-Origin': '*'

    })
}
 
  private SERVER = "/api/webapi/v1/user";
/*api functions*/
  get_userbyid(id){
  return  this.httpClient.get(this.SERVER+'/'+id);
  }

  post_data(userdata){
    return  this.httpClient.post(this.SERVER,userdata,this.httpOptions);
  }
  del_user(id){
    console.log("service",id);
    return this.httpClient.delete(this.SERVER+'/'+id,this.httpOptions);

  }
  update_user(id,userdata){
    return  this.httpClient.put(this.SERVER+'/'+id,userdata,this.httpOptions);
  }
  /*helper functions for interaction between two components */

  set_display_search(val){
   this.is_display_search=val; 
   console.log("set display",this.is_display_search);
  }
  get_display_search(){
    console.log("get display",this.is_display_search);
    return this.is_display_search;
  }
  set_display_update(val){
    this.is_display_update=val;
  }
  get_display_update(){
    return this.is_display_update;
  }
 
}

