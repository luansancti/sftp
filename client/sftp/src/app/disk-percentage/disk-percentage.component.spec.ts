import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DiskPercentageComponent } from './disk-percentage.component';

describe('DiskPercentageComponent', () => {
  let component: DiskPercentageComponent;
  let fixture: ComponentFixture<DiskPercentageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DiskPercentageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DiskPercentageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
