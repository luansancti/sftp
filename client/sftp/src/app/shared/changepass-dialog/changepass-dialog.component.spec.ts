import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ChangepassDialogComponent } from './changepass-dialog.component';

describe('ChangepassDialogComponent', () => {
  let component: ChangepassDialogComponent;
  let fixture: ComponentFixture<ChangepassDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ChangepassDialogComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ChangepassDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
