import { Injectable } from "@angular/core";

@Injectable({
    providedIn: 'root'
})
export class Helper {
    constructor() { }

    public formatBytes(bytes, decimals = 2): string {
        if (bytes === 0) return '0 Bytes';
    
        const k = 1024;
        const dm = decimals < 0 ? 0 : decimals;
        const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
    
        const i = Math.floor(Math.log(bytes) / Math.log(k));
    
        return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
    }

    public randomPass (length, addUpper, addSymbols, addNums) {
        var lower = "abcdefghijklmnopqrstuvwxyz";
        var upper = addUpper ? lower.toUpperCase() : "";
        var nums = addNums ? "0123456789" : "";
        var symbols = addSymbols ? "!#$%&\'()*+,-./:;<=>?@[\\]^_`{|}~" : "";
    
        var all = lower + upper + nums + symbols;
        while (true) {
            var pass = "";
            for (var i=0; i<length; i++) {
                pass += all[Math.random() * all.length | 0];
            }
    
            // criteria:
            if (!/[a-z]/.test(pass)) continue; // lowercase is a must
            if (addUpper && !/[A-Z]/.test(pass)) continue; // check uppercase
            if (addSymbols && !/\W/.test(pass)) continue; // check symbols
            if (addNums && !/\d/.test(pass)) continue; // check nums
    
            return pass; // all good
        }
    }
}
