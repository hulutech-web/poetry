export namespace main {
	
	export class Feather {
	    content: string;
	    origin: string;
	    author: string;
	    category: string;
	
	    static createFrom(source: any = {}) {
	        return new Feather(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.content = source["content"];
	        this.origin = source["origin"];
	        this.author = source["author"];
	        this.category = source["category"];
	    }
	}

}

