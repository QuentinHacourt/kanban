import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CreateStoryComponent } from './create-story/create-story.component';
import { KanbanComponent } from './kanban/kanban.component';
import { StoryDetailComponent } from './story-detail/story-detail.component';

const routes: Routes = [
  { path: '', redirectTo: '/kanban', pathMatch: 'full' },
  { path: 'kanban', component: KanbanComponent },
  { path: 'create-story', component: CreateStoryComponent },
  { path: 'story/:id', component: StoryDetailComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
