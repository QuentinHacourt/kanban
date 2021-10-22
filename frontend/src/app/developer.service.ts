import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Developer, DeveloperInput } from './developer';

@Injectable({
  providedIn: 'root',
})
export class DeveloperService {
  private developerUrl = 'http://localhost:8080/developer';
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
  };

  constructor(private http: HttpClient) {}

  getDevelopers(): Observable<Developer[]> {
    return this.http
      .get<Developer[]>(this.developerUrl)
      .pipe(catchError(this.handleError<Developer[]>('getDeveloper', [])));
  }

  getDeveloper(id: number): Observable<Developer> {
    const url = `${this.developerUrl}/${id}`;
    return this.http
      .get<Developer>(url)
      .pipe(catchError(this.handleError<Developer>(`getDeveloper id=${id}`)));
  }

  updateDeveloper(developer: Developer): Observable<any> {
    const url = `${this.developerUrl}/${developer.id}`;

    return this.http
      .put(url, developer, this.httpOptions)
      .pipe(catchError(this.handleError<any>('updateDeveloper')));
  }

  deleteDeveloper(id: number): Observable<Developer> {
    const url = `${this.developerUrl}/${id}`;

    return this.http
      .delete<Developer>(url, this.httpOptions)
      .pipe(catchError(this.handleError<Developer>('deleteDeveloper')));
  }

  addDeveloper(developer: DeveloperInput): Observable<Developer> {
    return this.http
      .post<Developer>(this.developerUrl, developer, this.httpOptions)
      .pipe(catchError(this.handleError<Developer>('addDeveloper')));
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
