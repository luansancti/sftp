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


class DiskUsage {
    DirectoryName: string;
    Percentage: number;
}


export class UserAdd {
    User: string;
    Expiration: number;
    Password: string;

}


export class UserDetails {
    UserName: string;
    Owner:  string;
    Key:    boolean;
    Expiration: any;
    Size: string;
}
