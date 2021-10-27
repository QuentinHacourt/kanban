import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Project, ProjectInput } from './project';

@Injectable({
  providedIn: 'root',
})
export class ProjectService {
  private projectUrl = 'http://localhost:8080/project';
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
  };

  constructor(private http: HttpClient) {}

  getProjects(): Observable<Project[]> {
    return this.http
      .get<Project[]>(this.projectUrl)
      .pipe(catchError(this.handleError<Project[]>('getProject', [])));
  }

  getProject(id: number): Observable<Project> {
    const url = `${this.projectUrl}/${id}`;
    return this.http
      .get<Project>(url)
      .pipe(catchError(this.handleError<Project>(`getProject id=${id}`)));
  }

  updateProject(project: Project): Observable<any> {
    const url = `${this.projectUrl}/${project.id}`;

    return this.http
      .put(url, project, this.httpOptions)
      .pipe(catchError(this.handleError<any>('updateProject')));
  }

  deleteProject(id: number): Observable<Project> {
    const url = `${this.projectUrl}/${id}`;

    return this.http
      .delete<Project>(url, this.httpOptions)
      .pipe(catchError(this.handleError<Project>('deleteProject')));
  }

  addProject(project: ProjectInput): Observable<Project> {
    return this.http
      .post<Project>(this.projectUrl, project, this.httpOptions)
      .pipe(catchError(this.handleError<Project>('addProject')));
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
