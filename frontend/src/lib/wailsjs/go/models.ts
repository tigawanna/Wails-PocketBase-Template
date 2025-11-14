export namespace main {
	
	export class ListItem {
	    id: string;
	    data: string;
	    state: boolean;
	    updated: string;
	
	    static createFrom(source: any = {}) {
	        return new ListItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.data = source["data"];
	        this.state = source["state"];
	        this.updated = source["updated"];
	    }
	}

}

