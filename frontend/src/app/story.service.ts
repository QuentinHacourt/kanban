import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { Story } from './story';
import { catchError, map, tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class StoryService {
  private storyUrl = 'http://localhost:8080/story';

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
  };

  constructor(private http: HttpClient) {}

  getStories(): Observable<Story[]> {
    return this.http
      .get<Story[]>(this.storyUrl)
      .pipe(catchError(this.handleError<Story[]>('getStory', [])));
  }

  getStory(id: number): Observable<Story> {
    const url = `${this.storyUrl}/${id}`;
    return this.http
      .get<Story>(url)
      .pipe(catchError(this.handleError<Story>(`getStory id=${id}`)));
  }

  updateStory(story: Story): Observable<any> {
    const url = `${this.storyUrl}/${story.id}`;

    return this.http
      .put(url, story, this.httpOptions)
      .pipe(catchError(this.handleError<any>('updateStory')));
  }

  deleteStory(id: number): Observable<Story> {
    const url = `${this.storyUrl}/${id}`;

    return this.http
      .delete<Story>(url, this.httpOptions)
      .pipe(catchError(this.handleError<Story>('deleteStory')));
  }

  addStory(story: Story): Observable<Story> {
    return this.http
      .post<Story>(this.storyUrl, story, this.httpOptions)
      .pipe(catchError(this.handleError<Story>('addStory')));
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
