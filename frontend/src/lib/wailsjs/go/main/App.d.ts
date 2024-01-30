// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {main} from '../models';

export function GetCurrentUser():Promise<main.User>;

export function GetReleaseInfo(arg1:boolean):Promise<main.ReleaseInfo>;

export function GetRequest(arg1:string):Promise<string>;

export function GetSettings():Promise<main.AppSettings>;

export function GetVersion():Promise<string>;

export function InstallUpdate():Promise<main.Response>;

export function PatchRequest(arg1:string,arg2:string):Promise<string>;

export function RefreshCurrentUser():Promise<string>;

export function SaveSettings(arg1:string):Promise<main.Response>;

export function StartAuthProcess():Promise<void>;
