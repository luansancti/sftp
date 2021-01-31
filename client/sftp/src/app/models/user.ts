import { Time } from '@angular/common';

export class ResponseUserList {
    Success: boolean;
    Message: string;
    Data: UserDetails[];
}

export class ResponseGeneric {
    Success: boolean;
    Message: string;
}

export class ResponseDiskPercentage {
    Success: boolean;
    Message: string;
    Data: DiskUsage[];
}

export class ResponseKey {
    Success: boolean;
    Message: string;
    Data: string;
}

export class ReponseListDirectory {
    Success: boolean;
    Message: string;
    Data: ListDirectory[];
}

export class ResponseData {
    Success: string;
    Data: string[];
}


export class DiskUsage {
    DirectoryName: string;
    Percentage: number;
}


export class UserAdd {
    User: string;
    Expiration: number;
    Password: string;

}

export class ListDirectory {
    Name: string;
    Size: number;
    IsDirectory: boolean;
    ModTime: Date;
}

export class UserDetails {
    UserName: string;
    Owner:  string;
    Key:    boolean;
    Expiration: any;
    Size: string;
}
