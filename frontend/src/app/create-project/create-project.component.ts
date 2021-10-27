import { Component, OnInit } from '@angular/core';
import { ProjectInput } from '../project';
import { ProjectService } from '../project.service';

@Component({
  selector: 'app-create-project',
  templateUrl: './create-project.component.html',
  styleUrls: ['./create-project.component.css'],
})
export class CreateProjectComponent implements OnInit {
  constructor(private projectService: ProjectService) {}

  ngOnInit(): void {}

  addProject(title: string, description: string): void {
    if (!title || !description) {
      return;
    }
    const projectInput: ProjectInput = {
      title: title,
      description: description,
    };
    this.projectService.addProject(projectInput).subscribe();
  }
}
