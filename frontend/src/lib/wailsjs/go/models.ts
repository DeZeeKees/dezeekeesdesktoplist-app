export namespace main {
	
	export class AppSettings {
	    usePrerelease: boolean;
	    yourListCardSizeMultiplier: number;
	
	    static createFrom(source: any = {}) {
	        return new AppSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.usePrerelease = source["usePrerelease"];
	        this.yourListCardSizeMultiplier = source["yourListCardSizeMultiplier"];
	    }
	}
	export class  {
	    name: string;
	    browser_download_url: string;
	
	    static createFrom(source: any = {}) {
	        return new (source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.browser_download_url = source["browser_download_url"];
	    }
	}
	export class Release {
	    assets: [];
	    tag_name: string;
	    prerelease: boolean;
	    published_at: string;
	    html_url: string;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new Release(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.assets = this.convertValues(source["assets"], );
	        this.tag_name = source["tag_name"];
	        this.prerelease = source["prerelease"];
	        this.published_at = source["published_at"];
	        this.html_url = source["html_url"];
	        this.message = source["message"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ReleaseInfo {
	    isLatest: boolean;
	    release: Release;
	
	    static createFrom(source: any = {}) {
	        return new ReleaseInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.isLatest = source["isLatest"];
	        this.release = this.convertValues(source["release"], Release);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class User {
	    id: number;
	    name: string;
	    picture: string;
	    gender: string;
	    birthday: string;
	    location: string;
	    joined_at: string;
	    // Go type: struct { NumItemsWatching int "json:\"num_items_watching\""; NumItemsCompleted int "json:\"num_items_completed\""; NumItemsOnHold int "json:\"num_items_on_hold\""; NumItemsDropped int "json:\"num_items_dropped\""; NumItemsPlanToWatch int "json:\"num_items_plan_to_watch\""; NumItems int "json:\"num_items\""; NumDaysWatched float64 "json:\"num_days_watched\""; NumDaysWatching float64 "json:\"num_days_watching\""; NumDaysCompleted float64 "json:\"num_days_completed\""; NumDaysOnHold float64 "json:\"num_days_on_hold\""; NumDaysDropped float64 "json:\"num_days_dropped\""; NumDays float64 "json:\"num_days\""; NumEpisodes int "json:\"num_episodes\""; NumTimesRewatched int "json:\"num_times_rewatched\""; MeanScore float64 "json:\"mean_score\"" }
	    anime_statistics: any;
	    time_zone: string;
	    is_supporter: boolean;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.picture = source["picture"];
	        this.gender = source["gender"];
	        this.birthday = source["birthday"];
	        this.location = source["location"];
	        this.joined_at = source["joined_at"];
	        this.anime_statistics = this.convertValues(source["anime_statistics"], Object);
	        this.time_zone = source["time_zone"];
	        this.is_supporter = source["is_supporter"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

