import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CreateProjectComponent } from './create-project/create-project.component';
import { CreateStoryComponent } from './create-story/create-story.component';
import { CreateTeamComponent } from './create-team/create-team.component';
import { DeveloperDetailsComponent } from './developer-details/developer-details.component';
import { DevelopersComponent } from './developers/developers.component';
import { KanbanComponent } from './kanban/kanban.component';
import { ProjectDetailsComponent } from './project-details/project-details.component';
import { ProjectsComponent } from './projects/projects.component';
import { RegisterDeveloperComponent } from './register-developer/register-developer.component';
import { StoryDetailComponent } from './story-detail/story-detail.component';
import { TeamDetailsComponent } from './team-details/team-details.component';
import { TeamsComponent } from './teams/teams.component';

const routes: Routes = [
  { path: '', redirectTo: '/kanban', pathMatch: 'full' },
  { path: 'kanban', component: KanbanComponent },
  { path: 'create-story', component: CreateStoryComponent },
  { path: 'story/:id', component: StoryDetailComponent },
  { path: 'developer', component: DevelopersComponent },
  { path: 'developer/register', component: RegisterDeveloperComponent },
  { path: 'developer/:id', component: DeveloperDetailsComponent },
  { path: 'project', component: ProjectsComponent },
  { path: 'project/create', component: CreateProjectComponent },
  { path: 'project/:id', component: ProjectDetailsComponent },
  { path: 'team', component: TeamsComponent },
  { path: 'team/create', component: CreateTeamComponent },
  { path: 'team/:id', component: TeamDetailsComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
