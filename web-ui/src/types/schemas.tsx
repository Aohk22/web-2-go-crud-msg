export interface User {
	Id: number, 
	Name: string 
};

export interface Message {
	Id: number,
	Time: string,
	Content: string,
	UserId: number,
}
