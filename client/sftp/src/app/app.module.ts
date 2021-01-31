import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { MatTableModule } from '@angular/material/table' 
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatSliderModule } from '@angular/material/slider';
import {MatPaginatorModule} from '@angular/material/paginator';
import {MatToolbarModule} from '@angular/material/toolbar';
import { MatSidenavModule } from '@angular/material/sidenav';
import {MatSortModule} from '@angular/material/sort';
import {MatExpansionModule} from '@angular/material/expansion';
import { DiskPercentageComponent } from './disk-percentage/disk-percentage.component';
import {MatInputModule } from '@angular/material/input';
import {MatIconModule} from '@angular/material/icon';
import { AddUserComponent } from './add-user/add-user.component';
import { MatPasswordStrengthModule } from '@angular-material-extensions/password-strength';
import {MatSlideToggleModule} from '@angular/material/slide-toggle';
import {FormsModule} from '@angular/forms'
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatButtonModule} from '@angular/material/button';
import {MatSelectModule} from '@angular/material/select';
import { ReactiveFormsModule } from '@angular/forms';
import {ClipboardModule} from '@angular/cdk/clipboard';
import {MatTooltipModule} from '@angular/material/tooltip';
import {MatSnackBarModule} from '@angular/material/snack-bar';
import {AddUserService} from './add-user/add-user.service'
import {MatDialogModule} from '@angular/material/dialog';
import { ConfirmationDialogComponent } from './shared/confirmation-dialog/confirmation-dialog.component';
import { ConfirmationDialogService } from './shared/confirmation-dialog/confirmation-dialog-service.service';
import { DashboardComponent } from './dashboard/dashboard.component';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import { ChangepassDialogComponent } from './shared/changepass-dialog/changepass-dialog.component';
import { RenewDialogComponent } from './shared/renew-dialog/renew-dialog.component';
import { RenewDialogService } from './shared/renew-dialog/renew-dialog.service';
import { TableDirectoryComponent } from './table-directory/table-directory.component';



@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    DiskPercentageComponent,
    AddUserComponent,
    ConfirmationDialogComponent,
    DashboardComponent,
    RenewDialogComponent,
    ChangepassDialogComponent,
    TableDirectoryComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    BrowserAnimationsModule,
    MatSliderModule,
    MatTableModule,
    MatPaginatorModule,
    MatToolbarModule,
    MatSidenavModule,
    MatSortModule,
    MatExpansionModule,
    MatInputModule,
    MatIconModule,
    MatPasswordStrengthModule.forRoot(),
    MatSlideToggleModule, 
    FormsModule,
    MatCheckboxModule,
    MatButtonModule,
    MatSelectModule,
    ReactiveFormsModule,
    ClipboardModule,
    MatTooltipModule,
    MatSnackBarModule,
    MatDialogModule,
    MatProgressSpinnerModule
  ],
  providers: [AddUserService, ConfirmationDialogService, RenewDialogService],
  bootstrap: [AppComponent],
  entryComponents: [MatDialogModule]
})
export class AppModule { }
