import { TestBed } from '@angular/core/testing';

import { RenewDialogService } from './renew-dialog.service';

describe('RenewDialogService', () => {
  let service: RenewDialogService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RenewDialogService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
