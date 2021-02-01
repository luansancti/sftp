import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomeComponent } from './home/home.component'
import {TableDirectoryComponent} from './table-directory/table-directory.component'

const routes: Routes = [
  { path: "", component: HomeComponent },
  { path: "userdirectory", component: TableDirectoryComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
