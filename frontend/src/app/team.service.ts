import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Team, TeamInput } from './team';

@Injectable({
  providedIn: 'root',
})
export class TeamService {
  private teamUrl = 'http://localhost:8080/team';
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
  };

  constructor(private http: HttpClient) {}

  getTeams(): Observable<Team[]> {
    return this.http
      .get<Team[]>(this.teamUrl)
      .pipe(catchError(this.handleError<Team[]>('getTeam', [])));
  }

  getTeam(id: number): Observable<Team> {
    const url = `${this.teamUrl}/${id}`;
    return this.http
      .get<Team>(url)
      .pipe(catchError(this.handleError<Team>(`getTeam id=${id}`)));
  }

  updateTeam(team: Team): Observable<any> {
    const url = `${this.teamUrl}/${team.id}`;

    return this.http
      .put(url, team, this.httpOptions)
      .pipe(catchError(this.handleError<any>('updateTeam')));
  }

  deleteTeam(id: number): Observable<Team> {
    const url = `${this.teamUrl}/${id}`;

    return this.http
      .delete<Team>(url, this.httpOptions)
      .pipe(catchError(this.handleError<Team>('deleteTeam')));
  }

  addTeam(team: TeamInput): Observable<Team> {
    return this.http
      .post<Team>(this.teamUrl, team, this.httpOptions)
      .pipe(catchError(this.handleError<Team>('addTeam')));
  }
  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {
      console.error(error); // log to console instead

      alert(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  /** Log a HeroService message with the MessageService */
  private log(message: string) {
    console.log(message);
  }
}
