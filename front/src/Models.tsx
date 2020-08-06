export class Facet {
  public val: string;
  public count: Number;
  public constructor(obj: any) {
    this.val = obj.val;
    this.count = obj.count;
  }
}
export class Article {
  public id: string;
  public url: string;
  public displayName: string;
  public userName: string;
  public spaceName: string;
  public title: string;
  public view: string;
  public constructor(obj: any) {
    this.id = obj.id;
    this.url = obj.URL;
    this.displayName = obj.createdBy_displayName;
    this.userName = obj.createdBy_username;
    this.spaceName = obj.space_name;
    this.title = obj.title;
    this.view = obj.view;
  }
}
