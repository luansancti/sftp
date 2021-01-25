import { TestBed } from '@angular/core/testing';

import { DiskPercentageService } from './disk-percentage.service';

describe('DiskPercentageService', () => {
  let service: DiskPercentageService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(DiskPercentageService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
