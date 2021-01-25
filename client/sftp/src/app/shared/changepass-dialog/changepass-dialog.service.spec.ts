import { TestBed } from '@angular/core/testing';

import { ChangepassDialogService } from './changepass-dialog.service';

describe('ChangepassDialogService', () => {
  let service: ChangepassDialogService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ChangepassDialogService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
