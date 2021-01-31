import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TableDirectoryComponent } from './table-directory.component';

describe('TableDirectoryComponent', () => {
  let component: TableDirectoryComponent;
  let fixture: ComponentFixture<TableDirectoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TableDirectoryComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TableDirectoryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
