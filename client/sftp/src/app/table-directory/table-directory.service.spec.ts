import { TestBed } from '@angular/core/testing';

import { TableDirectoryService } from './table-directory.service';

describe('TableDirectoryService', () => {
  let service: TableDirectoryService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TableDirectoryService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
